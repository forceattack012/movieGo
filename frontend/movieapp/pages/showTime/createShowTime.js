import { useState } from 'react'
import Layout  from '../../components/layout'
import Modal from '../../components/modal'
import { useForm, Controller } from "react-hook-form";
import Select from "react-select";
import moment from 'moment';

export default function CreateShowTime({theaters, movies}) {
    const [isModal, setModal] = useState(false);

    const { register, handleSubmit, control, formState : {errors}, } = useForm({
        mode: "onBlur"
    });

    const selectOptionsTheater = [];
    const selectOptionsMovie = [];

    theaters.map(theater => {
        selectOptionsTheater.push({value: theater.id, label: theater.name })
    });

    movies.map(movie => {
        selectOptionsMovie.push({value: movie.id, label: movie.name })
    });

    const createShowTime = async(form) => {
        const f = {
            time : moment(form.startDate),
            movieId : form.movie.value,
            theaterId : form.theater.value,
        }

        const create = JSON.stringify(f)

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: create
        };

        const response = await fetch('http://localhost:5001/api/v1/showtime/', requestOptions)

        if(response.status == 201) {
            setModal(true)
        }
    }

    return (
        <Layout>
                <Modal isModal={isModal} />
                <div className="md:container md:mx-auto">
                    <form className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" onSubmit={handleSubmit(createShowTime)}>
                        <div className="mb-4">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="theater">Theater</label>
                            <Controller
                                name="theater"
                                control={control}
                                defaultValue=""
                                render={({ field }) => (
                                    <Select key={field.value} options={selectOptionsTheater} {...field} label="Text field" />
                                )}
                            />
                        </div>

                        <div className="mb-4">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="theater">Movie</label>
                            <Controller
                                name="movie"
                                control={control}
                                defaultValue=""
                                render={({ field }) => (
                                    <Select key={field.value} options={selectOptionsMovie} {...field} label="Text field" />
                                )}
                            />
                        </div>

                        <div className="mb-4">
                            <label htmlFor="startDate">Start Date</label>
                            <div className="flex flex-wrap relative">
                                <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                                    <svg className="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clipRule="evenodd"></path></svg>
                                </div>
                                <input datepicker="true" type="datetime-local" id="startDate" name="startDate" className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date" 
                                {...register('startDate', {required: true})}/>
                                {errors.startDate && errors.startDate.type == 'required' && <p className='text-red-600'>Please Select Start Date</p>}
                            </div>
                        </div>

                        <div className="mb-4 flex justify-end">
                            <button type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
                            Create</button>
                        </div>
                    </form>
                </div>
        </Layout>
    )
}

export async function getStaticProps() {
    const resTheater = await fetch('http://localhost:5001/api/v1/theater/list')
    const resMovie = await fetch('http://localhost:5001/api/v1/movie/list')

    const theaters = await resTheater.json()
    const movies = await resMovie.json();

    return { props: { theaters, movies }}
}
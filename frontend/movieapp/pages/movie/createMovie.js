import {React, useState} from 'react';
import Layout from "../../components/layout";
import UploadFile from "../../components/uploadFile";
import { useForm, useFieldArray } from 'react-hook-form'
import moment from 'moment';
import Loading from '../../components/loading';
import Modal from '../../components/modal';



export default function CreateMovie() {
    const {register, control , handleSubmit , formState : {errors},} = useForm({
        defaultValues: {
            actors: [{name: ""}]
        }
    });

    const { fields, append, prepend, remove, swap, move, insert } = useFieldArray(
        {
          control,
          name: "actors"
        }
      );

    const [userImage, setUserImage] = useState(null);
    const [isLoading, setLoading] = useState(false);
    const [isModal, setModal] = useState(false);

    const imageUploadHandler = (event) => {
        setUserImage(event.target.files[0]);
    };

    const createMovie = async(form) => {
        setLoading(true)

        const f = {
            name: form.name,
            IMDB: form.IMDB,
            image: await getBase64(userImage).then(async data => {
                return await data
            }),
            startDate: moment(form.startDate, 'YYYY-MM-DD').toDate(),
            duration: parseInt(form.duration),
            youtube: form.youtube,
            actors:[],
            directors: [form.director],
            synopsis: form.synopsis
        }


        form.actors.forEach(item => {
            f.actors.push(item.name)
        })
        const create = JSON.stringify(f)
        console.log(create);

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: create
        };

        const response = await fetch('http://localhost:5001/api/v1/movie/', requestOptions)
        const result = response.status;
        if(result == 201){
            setLoading(false)
            setModal(true)
        }
    }

    const getBase64 = (file) => {
        return new Promise((resolve, reject) => {
          const reader = new FileReader();
          reader.readAsDataURL(file);
          reader.onload = () => resolve(reader.result);
          reader.onerror = error => reject(error);
        });
      }


    return (
        <Layout>
                <Modal isModal={isModal} />
                <div className="md:container md:mx-auto">
                    <form className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" encType="multipart/form-data" onSubmit={handleSubmit(createMovie)}>
                        <div className="mb-4">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="name">Name</label>
                            <input className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" type="text" 
                                {...register('name', {required: true})} />
                                {errors.name && errors.name.type == 'required' && <p className='text-red-600'>Please Enter Name of movie</p>}
                        </div>

                        <div className="mb-4">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="movieImage">Image</label>
                            <UploadFile imageUploadHandler={imageUploadHandler} image={userImage}/>
                        </div>

                        <div className="mb-3">
                            <div className='flex'>
                                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300 mr-2" htmlFor="actor">Actor</label>
                                <button className='mt-2 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none mb-2' type="button" 
                                    onClick={() => {insert(parseInt(2, 10), "actors" ,{ name : '' })}}>Add</button>
                            </div>
                            {
                                    fields.map((item, index) => {
                                        return (
                                            <div className="flex flex-wrap mb-2" key={item.id}>
                                                <div className='w-full md:w-5/6 mb-6 md:mb-0'>
                                                    <input 
                                                    name={`actors[${index}].name`}
                                                    className="mr-2 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    type="text"
                                                    {...register(`actors[${index}].name`)}
                                                    />
                                                </div>
                                                <div className='mb-3'>
                                                    <button className='mt-2 ml-3 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none' type="button" onClick={() => {remove(index)}}>Remove</button>
                                                </div>
                                                 <br></br>
                                            </div>
                                        )
                                    })
                                }
                        </div>

                        <div className="mb-3">
                                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300 mr-2" htmlFor="actor">Director</label>
                                <div className='w-full md:w-5/6 mb-6 md:mb-0'>
                                                    <input 
                                                    name="director"
                                                    className="mr-2 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    type="text"
                                                    {...register("director")}
                                                    />
                                </div>
                        </div>

                        <div className="mb-4">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="imdb">IMDB</label>
                            <input className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            type="text" id="imdb" name="imdb" {...register('IMDB', {pattern:/^([1-9]|10)\.[0-9][0-9]|[1-7]$ /} )}/>
                            {errors.IMDB && errors.IMDB.type == 'pattern' && <p className='text-red-600'>Please Enter IMDB only number 1 - 10</p>}
                        </div>

                        <div className="flex flex-wrap -mx-3 mb-2">
                            <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="youtube">Youtube</label>
                                <input className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                type="text" id="youtube" name="youtube" {...register('youtube', {pattern:/^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/})}/>
                                {errors.youtube && errors.youtube.type == 'pattern' && <p className='text-red-600'>Please Enter youtube URL</p>} 
                            </div>

                            <div className="w-full md:w-1/2 px-3">
                                <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="duration">Duration</label>
                                <input className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                type="text" id="duration" name="duration" {...register('duration',{required: true, pattern: /\d+/ })} />
                                { errors.duration && 
                                   <div>
                                        {errors.duration.type == 'required' && <p className='text-red-600'>Please Enter duration</p> }
                                        {errors.duration.type == 'pattern' && <p className='text-red-600'>Please Enter duration number only</p>}
                                   </div> 
                                }
                            </div>
                        </div>

                        <div className="mb-4">
                            <label htmlFor="startDate">Start Date</label>
                            <div className="flex flex-wrap relative">
                                <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                                    <svg className="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clipRule="evenodd"></path></svg>
                                </div>
                                <input datepicker="true" type="date" id="startDate" name="startDate" className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date" 
                                {...register('startDate', {required: true})}/>
                                {errors.startDate && errors.startDate.type == 'required' && <p className='text-red-600'>Please Select Start Date</p>}
                            </div>
                        </div>

                        <div className="mb-4">  
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="synopsis">Synopsis</label>
                            <textarea id="synopsis" name="synopsis" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
                            {...register('synopsis')}></textarea>
                        </div>

                        <div className="mb-4 flex justify-end">
                            <button type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
                            Create</button>
                        </div>
                    </form>
                </div>
                <Loading isLoading={isLoading} />
        </Layout>
    )
}
import Layout from "../../components/layout";
import {useForm} from 'react-hook-form'
import { useState } from "react";
import Modal from '../../components/modal';

export default function CreateTheater() {

    const {register, control , handleSubmit , formState : {errors},} = useForm({
        defaultValues: {
            isOpen: [true]
        }
    })

    const [isModal, setModal] = useState(false)

    const createTheater = async(form) => {
        const f = {
            name: form.theaterName,
            isOpen: form.isOpen == "true" ? true : false
        }

        const create = JSON.stringify(f)
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: create
        };

        const response = await fetch('http://localhost:5001/api/v1/theater/', requestOptions)

        if(response.status == 201) {
            setModal(true)
        }

    }

    return (
        <Layout>
                <Modal isModal={isModal} />
                <div className="md:container md:mx-auto">
                    <form className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" encType="multipart/form-data" onSubmit={handleSubmit(createTheater)}>
                        <div className="mb-4">
                            <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" htmlFor="name">Theater Name</label>
                            <input className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" type="text" 
                                {...register('theaterName', {required: true})} />
                                {errors.name && errors.name.type == 'required' && <p className='text-red-600'>Please Enter Name of theater</p>}
                            
                        </div>

                        <div className="flex items-start mb-6">
                                <div className="flex items-center h-5">
                                    <input aria-describedby="open" type="checkbox" className="w-4 h-4 bg-gray-50 rounded border border-gray-300 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800"
                                            {...register('isOpen')} defaultValue={true} />
                                </div>
                                <div className="ml-3 text-sm">
                                    <label htmlFor="Open" className="font-medium text-gray-900 dark:text-gray-300">Open</label>
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
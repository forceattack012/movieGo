import { useState } from "react";
import Layout from "../components/layout";
import Loading from "../components/loading";
import { useRouter } from "next/router"

export default function Movie({movies}) {
    const [isLoading, setLoading] = useState(false);
    const router = useRouter()

    const deleteMovie = async (id) => {
        const requestOptions = {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' }
        };

        const res = await fetch(`http://localhost:5001/api/v1/movie/${id}`, requestOptions)
        const json = res.json();

        if(res.status == 200 && json){
            router.reload()
        }
    }

    const editMove = async(id) => {
        router.push(`movie/edit/${id}`)
    }

    const gotoBuy = async(id) => {
        router.push(`/showTime/view//${id}`)
    }

    return (
        <Layout>
            <Loading isLoading={isLoading} />
            <h1 className="ml-20 text-3xl md:text-4x1 font-medium mb-2 text-left">
                Now Showing
            </h1>
            <div className="flex justify-center">
                <div className="mt-5 inline-grid grid-cols-4 gap-4">
                    {
                        movies.map(movie => {
                            return (
                                movie.id != '' &&                             
                                <div className="flex justify-center" key={movie.id.toString()}>
                                    <div className="rounded-lg shadow-lg bg-white max-w-sm hover:bg-gray-200">
                                        <a href={`/movie/view/${movie.id.toString()}`}>
                                        {movie.image != null && <img className="rounded-t-lg" src={movie.image} style={{width: "400px", height: "400px"}} alt=""></img>}
                                        {movie?.image == '' && <img className="rounded-t-lg" src="https://mdbootstrap.com/img/new/standard/nature/184.jpg" alt=""/>}
                                        </a>
                                        <div className="p-6">
                                            <h5 className="text-gray-900 text-xl font-medium mb-2">{movie.name}</h5>
                                            <p className="break-all text-gray-700 text-base mb-4">
                                                {movie.synopsis}
                                            </p>
                                            <button type="button" onClick={() => editMove(movie.id)} className="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Edit</button>
                                            <button type="button" onClick={() => gotoBuy(movie.id)} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Buy</button>
                                            <button type="button" onClick={() => deleteMovie(movie.id)} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-red-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-red-700 hover:shadow-lg focus:bg-red-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-red-800 active:shadow-lg transition duration-150 ease-in-out">Delete</button>
                                        </div>
                                    </div>
                                </div>
                            )
                        })
                    }
                </div>
            </div>

        </Layout>
    )
}

export async function getStaticProps() {
    const response = await fetch('http://localhost:5001/api/v1/movie/list')
    const movies = await response.json();

    return { props: { movies }}
}
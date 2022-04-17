import { useState } from "react";
import Layout from "../../../components/layout";
import Loading from "../../../components/loading";
import moment from 'moment';
import { useRouter } from "next/router"

export default function ViewShowTimes({showTimes, movie}) {
    const [isLoading, setLoading] = useState(false)
    const router = useRouter()

    const deleteShowTime = async(id) => {
        const requestOptions = {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' }
        };

        const res = await fetch(`http://localhost:5001/api/v1/showtime/remove/${id}`, requestOptions)
        const json = res.json();

        if(res.status == 404 && json){
            router.reload()
        }
    }

    console.log(showTimes)
    return (
        <Layout>
            <Loading isLoading={isLoading} />
            <h1 className="ml-20 text-3xl md:text-4x1 font-medium mb-2 text-left">
                Show Times
            </h1>
            {
                showTimes.map(showtime => {
                    return (
                        <div key={showtime.id}>
                            <div className="ml-20 mb-5 flex flex-col items-center bg-white rounded-lg border shadow-md md:flex-row hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 max-w-[90%]">
                                {movie.image != null && <img className="object-cover w-full h-96 rounded-t-lg md:h-auto md:w-48 md:rounded-none md:rounded-l-lg rounded-t-lg" src={movie.image} style={{width: "200px", height: "200px"}} alt=""></img>}

                                <div className="flex flex-col justify-between p-4 leading-normal">
                                    <div className="text-gray-900 text-xl font-medium mb-2">{showtime.Theater.name}</div>
                                        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">{movie.synopsis}</p>
                                    <div>
                                            {showtime.times.map((time, index) => {
                                                return (
                                                    <button key={index} type="button" onClick={() => console.log(showtime.id)} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">
                                                        {moment(time).format("HH:mm")}
                                                    </button>
                                                )
                                            })}
                                        <button type="button" onClick={() => editMove(showtime.id)} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">
                                            edit
                                        </button>
                                        <button type="button" onClick={() => deleteShowTime(showtime.id)} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-red-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-red-700 hover:shadow-lg focus:bg-red-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-red-800 active:shadow-lg transition duration-150 ease-in-out">
                                            Delete
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    )
                })
            }

        </Layout>
    )
}

export async function getStaticPaths() {
    const result = await fetch('http://localhost:5001/api/v1/showtime/list')
    const showTimes = await result.json()

    const paths = showTimes.map(showTime => {
        return {
            params : {
                movieId : `${showTime.movieId}`  // movieId same file name
            }
        }
    })

    return {
        paths,
        fallback: true // false or 'blocking'
      };
}

export async function getStaticProps(context) {
    const { params } = context;

    const req = await fetch(`http://localhost:5001/api/v1/showtime/${params.movieId}`)
    const showTimes = await req.json()

    const result = await fetch(`http://localhost:5001/api/v1/movie/${params.movieId}`)
    const movie = await result.json()

    return {
        props : {showTimes, movie}
    }
}
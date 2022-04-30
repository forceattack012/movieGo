import Layout from "../../components/layout";
import { useState } from "react";
import moment from 'moment';

export default function Booking({showTimes}) {
    const [selectSeats, setSeat] = useState('')
    const [total, setTotal] = useState(0)

    console.log(showTimes);

    const submit = async() =>{
        const result = [];
        const split = selectSeats.split(' ');
        split.map(s => {
            const filter = showTimes.seats.filter(st => st.seatNumber == s)
            filter.forEach(element => {
                result.push({
                    seatNumber: element.seatNumber,
                    theaterId: element.theaterId,
                    price: element.price
                })
            });
            return result
        })

        const booking = {
            showtimeId: showTimes.showTimeId,
            timestamp: moment(Date.UTC, 'YYYY-MM-DD').toDate(),
            seat: result.filter(s => s.length != 0),
            total: total
        } 

        const create = JSON.stringify(booking)
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: create
        };

        const response = await fetch('http://localhost:5001/api/v1/booking/', requestOptions)
    }
    return (
        <Layout>
            <h1 className="ml-20 text-3xl md:text-4x1 font-medium mb-2 text-left">
             Booking
            </h1>
            <div className="flex justify-center">
                <div className="mt-5 inline-grid grid-cols-4 gap-4">
                {
                showTimes.seats.map(seat=> {
                    return (
                        showTimes.bookingSeats.length > 0 && 
                            showTimes.bookingSeats.map(bk => {
                                if(seat.id === bk.id) {
                                    return 'sdsdsdsadas'
                                }else {
                                   return <div key={seat.id}>
                                        <button key={seat.seatNumber} onClick={ () => {setSeat(selectSeats +' '+ showTimes.seats.filter(s => s.seatNumber == seat.seatNumber).seatNumber), setTotal(total + seat.price) }} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">{seat.seatNumber}</button>
                                   </div>
                                }
                            })
                    )
                })
            }

            {
                showTimes.seats.map(seat=> {
                    return (
                        showTimes.bookingSeats.length == 0 && 
                        <div key={seat.id}>
                              <button key={seat.seatNumber} onClick={ () => {setSeat(selectSeats + ' '+ seat.seatNumber), setTotal(total + seat.price) }} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">{seat.seatNumber}</button>
                        </div>
                    )
                })
            }
                </div>
            </div>
            <br></br>
            <br></br>
            <div key={showTimes.id}>
                            <div className="ml-20 mb-5 flex flex-col items-center bg-white rounded-lg border shadow-md md:flex-row hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 max-w-[90%]">
                                {showTimes.image != null && <img className="object-cover w-full h-96 rounded-t-lg md:h-auto md:w-48 md:rounded-none md:rounded-l-lg rounded-t-lg" src={showTimes.image} style={{width: "200px", height: "200px"}} alt=""></img>}

                                <div className="flex flex-col justify-between p-4 leading-normal">
                                    <div className="text-gray-900 text-xl font-medium mb-2">{showTimes.movieName}</div>
                                        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">{showTimes.theaterName}</p>
                                        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">Time : {showTimes.duration}</p>
                                        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">Seats: {selectSeats}</p>
                                        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">Total : {total}</p>

                                        <button type="button" onClick={() => submit()} className="mr-3 ml-2 inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">
                                            Submit
                                        </button>
                                </div>
                                
                            </div>
                            
            </div>
        </Layout>
    )
}


export async function getStaticPaths() {
    const result = await fetch('http://localhost:5001/api/v1/theater/list')
    const showTimes = await result.json()

    const paths = showTimes.map(showTime => {
        return {
            params : {
                id : `${showTime.id}`  // movieId same file name
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

    const req = await fetch(`http://localhost:5001/api/v1/booking/${params.id}`)
    const showTimes = await req.json()

    return {
        props : {showTimes}
    }
}
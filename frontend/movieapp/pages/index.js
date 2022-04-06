import { useRouter } from 'next/router'
import { useEffect } from 'react'

export default function Home({movies}) {

  const route = useRouter()

  useEffect(() => {
    setTimeout(() => {
      route.push('/movie')
    }, 1000);

  },[])


  return (
    <><script src="https://unpkg.com/flowbite@1.4.1/dist/datepicker.js"></script><div>Welcome MovieGo</div></>
  )
}




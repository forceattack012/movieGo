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
    <div>Welcome MovieGo</div>
  )
}




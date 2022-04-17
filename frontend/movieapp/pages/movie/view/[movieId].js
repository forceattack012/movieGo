import Layout  from '../../../components/layout'

export default function ViewMovie({movie}) {
    return (
        <Layout>
            <div>{movie.id}</div>
            <div>{movie.name}</div>
        </Layout>
    )
}

export async function getStaticPaths() {
    const res = await fetch('http://localhost:5001/api/v1/movie/list')
    const movies = await res.json()

    const paths = movies.map(movie => {
        return {
            params : {
                movieId : `${movie.id.toString()}`  // movieId same file name
            }
        }
    })

    return {
        paths,
        fallback: true // false or 'blocking'
      };
    }

export async function getStaticProps(context) {
    const { params } = context

    const res = await fetch(`http://localhost:5001/api/v1/movie/${params.movieId}`)
    const movie = await res.json()

    return {
        props : {movie}
    }
}
import Layout from "../components/layout";

export default function Movie({movies}) {
    return (
        <Layout>
            <h1 className="text-3xl md:text-4x1 font-medium mb-2 text-left">
                Now Showing
            </h1>
            {
                movies.map(movie => {
                    return (
                        <div key={movie.Id}>
                            {movie.Name}
                        </div>
                    )
                })
            }

        </Layout>
    )
}

export async function getStaticProps() {
    const response = await fetch('http://localhost:5001/api/v1/movie/list')
    const movies = await response.json()
    console.log(movies);

    return { props: { movies  }}
}
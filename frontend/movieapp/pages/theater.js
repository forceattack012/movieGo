import Layout from "../components/layout";


export default function Theater({theaters}) {
    return (
        <Layout>
            {
                theaters.map(theater => {
                    return (
                        <div key={theater.id}>{theater.name}</div>
                    )
                })
            }
        </Layout>
    )
}

export async function getStaticProps() {
    const req = await fetch('http://localhost:5001/api/v1/theater/list');
    const theaters = await req.json();
    console.log(theaters);

    return { props: { theaters }}
}
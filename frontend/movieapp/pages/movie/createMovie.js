import Layout from "../../components/layout";

export default function CreateMovie() {
    const createMovie = async(event) => {
        event.preventDefault();
        console.log(event.target.name.value)
    }

    return (
        <Layout>
            <div>
                <form onSubmit={createMovie}>
                    <label htmlFor="name">Name</label>
                    <input type="text" id="name" name="name" />

                    <label htmlFor="actor">Actor</label>
                    <input type="text" id="actor" name="actor" />

                    <label htmlFor="imdb">IMDB</label>
                    <input type="text" id="imdb" name="imdb"/>

                    <label htmlFor="youtube">Youtube</label>
                    <input type="text" id="youtube" name="youtube"/>

                    <label htmlFor="duration">Duration</label>
                    <input type="text" id="duration" name="duration"></input>

                    <label htmlFor="startDate">Start Date</label>
                    <input type="date" id="startDate" name="startDate"></input>

                    <label htmlFor="synopsis">Synopsis</label>
                    <input type="area" id="synopsis" name="synopsis"></input>

                    <button type="submit">Create</button>
                </form>
            </div>
        </Layout>
    )
}
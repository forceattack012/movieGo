import Header from "./header";

export default function Layout(props) {
    return (
        <div>
            <Header />
            <div className="bg-indigo-100 py-6 md:py-12">
                {props.children}
            </div>
        </div>
    )
}
import LinkComponent from './link'
import Link from 'next/link';


export default function Header() {
    return (
        <header>
            <nav className="bg-white py-2 md:py-4">
                <div className="container px-4 mx-auto md:flex md:items-center">
                    <div className="flex justify-between items-center">
                        <Link href="/">
                            <a className="font-bold text-xl text-indigo-600">
                                MovieGo
                            </a>
                        </Link>
                        <button className="border border-solid border-gray-600 px-3 py-1 rounded text-gray-600 opacity-50 hover:opacity-75 md:hidden" id="navbar-toggle">
                            <i className="fas fa-bars"></i>
                        </button>
                    </div>

                    <div className="hidden md:flex flex-col md:flex-row md:ml-auto mt-3 md:mt-0">
                        <LinkComponent path={"/"} name="">home</LinkComponent>
                        <LinkComponent path={"/"} name="movie">movie</LinkComponent>
                        <LinkComponent path={"/"} name="theater">theater</LinkComponent>
                        <LinkComponent path={"/movie"} name="/createMovie">Create Movie</LinkComponent>
                        <LinkComponent path={"/theater"} name="/createTheater">Create Theater</LinkComponent>
                        <LinkComponent path={"/showTime"} name="/createShowTime">Create Show time</LinkComponent>
                    </div>
                </div>
            </nav>
        </header>
    )
}
import Link from 'next/link';
import { useRouter } from 'next/router';

export default function LinkComponent(props) {
    const router = new useRouter();
    const name = props.children;
    return (
        name == "home" ? 
        (
            <Link href="/">
                <a className={`p-2 lg:px-4 md:mx-2 uppercase
                    ${router.pathname == "/" ? "rounded bg-indigo-600 text-white" : "text-gray-600 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300" }`}>
                    {name}
                </a>
            </Link>
        ) : 
        (
            <Link href={`/${name}`}>
                <a className={`p-2 lg:px-4 md:mx-2 uppercase
                    ${router.pathname == `/${name}`  ? "rounded bg-indigo-600 text-white" : "text-gray-600 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300" }`}>
                    {name}
                </a>
            </Link>
        )
    )
}
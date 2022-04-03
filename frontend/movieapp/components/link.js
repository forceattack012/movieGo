import Link from 'next/link';
import { useRouter } from 'next/router';

export default function LinkComponent(props) {
    const router = new useRouter()
    const path = props.path
    const name = props.name

    return (
        <Link href={`${path}${name}`}>
            <a className={`p-2 lg:px-4 md:mx-2 uppercase
                ${router.pathname == `${path}${name}`  ? "rounded bg-indigo-600 text-white" : "text-gray-600 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300" }`}>
                {props.children}
            </a>
        </Link>
    )
}
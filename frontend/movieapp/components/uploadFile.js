
export default function UploadFile({ imageUploadHandler, image }) {
    return(
        <div>
            <label type="submit" className="mt-2 rounded-sm px-3 py-1 bg-gray-200 hover:bg-gray-300 focus:shadow-outline focus:outline-none">
                Upload a file
                <input id="uploadFile" type="file" name="fileUpload" onChange={imageUploadHandler} accept="image/*" className="opacity-0"
                multiple hidden/>
            </label>
            <div>
                {image && <img src={URL.createObjectURL(image)} style={{width: "400px"}} alt="preview"></img>}
            </div>
        </div>
    )
}
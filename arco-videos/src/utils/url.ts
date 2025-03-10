export const UrlParse = (
    url:string
)=>{
    if (url === undefined || url === "") return "";
    console.log()
    const s =url.match(/\d+/gim)
    if(!s) return "-";
    return s[0]
}

export default "-";
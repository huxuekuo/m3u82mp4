export const UrlParse = (
    url:string
)=>{
    if (url === undefined || url === "") return "";
    const s =url.match(/\d+/gim)
    console.log(s)
    if(!s) return "-";
    return s[s.length-1]
}

export default "-";
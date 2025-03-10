export const UrlParse = (
    url:string
)=>{
    const s =url.match(/\d+/gim)
    if(!s) return "-";
    return s[0]
}

export default "-";
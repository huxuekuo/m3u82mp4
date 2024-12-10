export const UrlParse = (
    url:string
)=>{
    const s =url.match(/\d+/gim)
    if(!s) return "-";
    console.log(s[0])
    return s[0]
}

export default "-";
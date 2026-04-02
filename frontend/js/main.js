import { RenderRegister }  from "./handlers/register.js";
import { RenderLogin } from "./handlers/login.js";  
import { Router } from "./handlers/router.js";
RenderRegister()


window.addEventListener("popstate", ()=> {
    Router()
})
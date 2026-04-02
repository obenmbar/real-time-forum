 
 import { RenderRegister } from "./register";
 import { RenderLogin } from "./login";
import { RenderHome } from "./home";

  export const Router = ()=> {
        const path = window.location.pathname
        if (path === "/register"){
            RenderRegister()
        }else if (path === "/login") {
         RenderLogin()
        }else {
            RenderHome()
        }
 }

 export function Navigate(path){ 
    window.history.pushState({},"",path)

    Router()
 }
 
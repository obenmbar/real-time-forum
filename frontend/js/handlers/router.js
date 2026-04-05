 
 import { RenderRegister } from "./register.js";
 import { RenderLogin } from "./login.js";
import { RenderHome } from "./home.js";

  export const Router = ()=> {
        const path = window.location.hash
        if (path === "#/register"){
            RenderRegister()
        }else if (path === "#/login") {
         RenderLogin()
        }else {
            RenderHome()
        }
 }

 export function Navigate(path){ 
    window.location.hash =  path
 }
 
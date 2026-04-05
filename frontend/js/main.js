import { RenderHome } from "./handlers/home.js"
import { Router } from "./handlers/router.js"

RenderHome()


window.addEventListener("popstate", ()=> {
  Router()
})

 window.addEventListener('hashchange',()=> {
    Router()
 })

 window.addEventListener("load",()=> {
    if (!window.location.hash || window.location.hash === "#"){
        window.location.hash ="/"
    }
 })

 localStorage.setItem("logout",Date.now())
 localStorage.addEventListener("storage",(event)=> {
         if (event.key === "logut") {
            window.location.reload()
         }
 })
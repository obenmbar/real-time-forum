
import { RenderRegister } from "./register.js"
import { RenderHome } from "./home.js"
import { Navigate } from "./router.js"

export function RenderLogin(){
    const container =   document.getElementById("container")
     const loginhtmlpage =`<form id="login-form">
        <h2>Login</h2>
        
        <div class="input-group">
            <label for="user">Email or Username:</label>
            <input type="text" name="user" id="user" required placeholder="Enter your Email or Username">
        </div>

        <div class="input-group">
            <label for="password">Password:</label>
            <input type="password" name="password" id="password" required placeholder="Enter your password" maxlength="30" minlength="8">
        </div>

        <button type="submit" id="submit-btn">Submit</button>

        <div id="container-linkregister">
            <p>Are you a new member? <a href="#" id="to-register">Register</a></p> 
        </div>
    </form>`

    container.innerHTML = loginhtmlpage
    
    LoginLogique()

}


export function LoginLogique() {
   const form = document.getElementById('login-form')
    const submit_btn =  document.getElementById('submit-btn')

   form.addEventListener("submit", async (event)=> {
    event.preventDefault()
    submit_btn.disabled = true
    submit_btn.innerText = "Logging in..."
    const formdata =  new FormData(form)
    const data =  Object.fromEntries(formdata.entries())

   if (data.password.length<8 || data.password.length >30){
        alert("Password must be between 8 and 30 characters")
        return
     }
      
    const value = Object.values(data)
    if (value.includes('')){
      alert("all feileds are required!")
        return
    }
    
    try {
        const response = await fetch("/api/login", {
            method:"POST",
         headers:{
            "content-Type": "application/json"
         },
            body: JSON.stringify(data)    
        })

            const result =  await response.json()

        if (!response.ok){
           
            throw new Error(result.messege||"Login failed")
        }

        RenderHome()
        console.log(result);
        

    }catch(error){
          alert(error.message)
          return
    } finally {
        submit_btn.disabled = false;
        submit_btn.innerText = "submit"
    }

   })

    const linkregister = document.getElementById('to-register') 
    
    linkregister.addEventListener('click', (event)=> {
        event.preventDefault()
         Navigate("/register")
    })
}

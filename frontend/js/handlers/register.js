import { RenderHome } from "./home.js"
import { RenderLogin } from "./login.js"
import { Navigate } from "./router.js"

  export function RenderRegister() {
  const container = document.getElementById('container')
     
   const registerhtml = `<h2>Register</h2>
        <form id="register-form">
            <div class="input-group">
                <label for="nickname">Nickname</label>
                <input type="text" name="nickname" id="nickname" required placeholder="Enter nickname">
            </div>

            <div class="input-group">
                <label for="first_name">First Name</label>
                <input type="text" name="first_name" id="first_name" required placeholder="Enter first name">
            </div>

            <div class="input-group">
                <label for="last_name">Last Name</label>
                <input type="text" name="last_name" id="last_name" required placeholder="Enter last name">
            </div>

            <div class="input-group">
                <label for="age">Age</label>
                <input type="number" name="age" id="age" required placeholder="Enter age" max="100" min= "0">
            </div>

            <div class="input-group">
                <label for="gender">Gender</label>
                <select name="gender" id="gender" required>
                    <option value="" disabled selected>Select Gender</option>
                    <option value="Male">Male</option>
                    <option value="Female">Female</option>
                </select>
            </div>

            <div class="input-group">
                <label for="email">Email</label>
                <input type="email" name="email" id="email" required placeholder="Enter email">
            </div>

            <div class="input-group">
                <label for="password">Password</label>
                <input type="password" name="password" id="password" required placeholder="Enter password">
            </div>

            <div class="input-group">
                <label for="confirm-password">Confirm Password</label>
                <input type="password" name="confirm_password" id="confirm-password" required placeholder="Confirm password">
            </div>

            <button type="submit" id="submit-btn">Submit</button>

            <div id="container-linklogin">
            <p>Already have an account? <a href="#" id="to-login">Login</a></p> 
              </div>
        </form>`

    container.innerHTML= registerhtml

    RegisterLogique()
    
  }
 export const RegisterLogique = () => {
   

    const form  = document.getElementById('register-form')

    form.addEventListener('submit', async (event) => {
     event.preventDefault()
     const formData =  new FormData(form)
     const data = Object.fromEntries(formData.entries())
     if (data.password.length<8 || data.password.length >30){
        alert("valid password length enter 8 and 30 caracter")
        return
     }

      if (data.password !== data["confirm_password"]){
        alert('password not match')
        return
      }

       const values =  Object.values(data)
       if (values.includes("")){
        alert("all feileds are required!")
        return
       }
            
       const age = Number(data.age)
       if (isNaN(age)|| age < 10|| age >130) {

        alert("virefier l'age")
            return
       }

     const emailregex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/      
     if (!emailregex.test(data.email)){
        alert('invalid email')
        return 
     }

     try {
        let  response = await fetch('/api/register', {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "content-Type": "application/json"
            }
        })

        if (!response.ok){
            const errorresponse = await response.json()
            throw new Error(errorresponse.messege|| "somthing bad")
        }
        
     const result = await response.json()
     RenderHome()
      console.log(result)
    
    }catch(error){
       alert(error.message || "error en backend")
       console.error(error.message)
     }
     
    })
    const linklogin = document.getElementById('to-login')
    linklogin.addEventListener('click', (event)=> {
        event.preventDefault()
        Navigate("/login")
    })
 }

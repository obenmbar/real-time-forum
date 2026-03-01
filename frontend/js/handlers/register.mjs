


 export const initialregisterDAta = () => {
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
                <input type="number" name="age" id="age" required placeholder="Enter age">
            </div>

            <div class="input-group">
                <label for="gender">Gender</label>
                <select name="gender" id="gender" required>
                    <option value="" disabled selected>Select Gender</option>
                    <option value="male">Male</option>
                    <option value="female">Female</option>
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
                <input type="password" name="confirm-password" id="confirm-password" required placeholder="Confirm password">
            </div>

            <button type="submit" id="submit-btn">Submit</button>
        </form>`

    container.innerHTML= registerhtml

    const form  = document.getElementById('register-form')

    form.addEventListener('submit', async (event) => {
     event.preventDefault()
     const formData =  new FormData(form)
     const data = Object.fromEntries(formData.entries())
     if (data.password.length<6){
        alert("password lnegth small")
        return
     }

      if (data.password !== data["confirm-password"]){
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
        let  response = await fetch('/register', {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "content-Type": "application/json"
            }
        })

        if (!response.ok){
            const errorresponse = await response.json()
            throw new Error(errorresponse || "somthing bad")
        }
        
     const result = await response.json()
      console.log(result)
     }catch(error){
       alert(error.messege || "error en backend")
       console.error(error.messege)
     }
    
     
    })
 }

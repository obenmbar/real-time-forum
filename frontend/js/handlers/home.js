/**
 * The RenderHome function sets up the layout for a home page with a navigation sidebar, a feed section
 * for creating posts, and a messenger sidebar with user contacts.
 */
export function RenderHome() {
    const container = document.getElementById('container');

    // أول حاجة: حيد الـ Styles القديمة ديال الـ Login من الـ container
    container.style.maxWidth = "none";
    container.style.padding = "0";
    container.style.width = "100%";

    container.innerHTML = `
    <div class="home-layout">
        <aside class="sidebar-left">
            <nav class="main-nav">
                <div class="logo">Forum<span>Node</span></div>
                <ul>
                    <li class="active"><i class="fas fa-home"></i> Home</li>
                    <li><i class="fas fa-user"></i> Profile</li>
                    <li id="logout-btn"><i class="fas fa-sign-out-alt"></i> Logout</li>
                </ul>
            </nav>
        </aside>

        <main class="feed-section">
         <section class="create-post">
    <form id="post-form">
        <div class="input-group-post">
            <input 
                type="text" 
                id="post-title" 
                name="post-title" 
                placeholder="Post Title..." 
                required 
                minlength="5"
                maxlength="100">
        </div>

        <textarea 
            id="post-content" 
            name="post-content" 
            placeholder="What's on your mind, Othmane?" 
            required></textarea>

        <div class="post-actions">
            <select id="post-category" name="post-category">
                <option value="general" selected>🌍 General Discussion</option>
                <option value="programming">💻 Programming & Tech</option>
                <option value="gaming">🎮 Gaming</option>
                <option value="help">🆘 Help & Support</option>
                <option value="offtopic">☕ Off-Topic</option>
            </select>
            <button type="submit" class="btn-primary">Post</button>
        </div>
    </form>
</section>

            <section id="posts-container" class="posts-list">
                <div class="loader">Loading posts...</div>
            </section>
        </main>

        <aside class="messenger-sidebar">
            <header class="messenger-header">
                <h3>Contacts</h3>
            </header>
            <div id="user-search">
                <input type="text" id="search-user" name="search-user" placeholder="Search users...">
            </div>
            <ul id="users-list" class="users-list">
                <li class="user-item-loading">Loading users...</li>
            </ul>
        </aside>
    </div>
    `;


    // هنا من بعد غاتعيط للدالات ديالك:
    // InitWebSocket();
    // FetchAndRenderPosts();
    // FetchAndRenderUsers();
}

let socket;
export function initWebSocket() {
    socket = new WebSocket("ws://localhost:8050/ws")


    socket.onopen = () => {
        console.log("connection open");

        const msg = {
            type: "identify",
            user_id: "othmane1",
            content: "hello i'm othmane",
        }
        socket.send(JSON.stringify(msg))
    }

     socket.onmessage = (evnet)=> {
        console.log(event.data)
        const data = JSON.parse(event.data)
        if (data.type === "new messege") {
            Rendermessege(data)
        }else if (data.type === "user_status") {
            updateUserStatus(data);
        }else if (data.type === "typing"){
            RenderTyping(data)
        }else if (data.type === "error"){
            alert("data.messege")
        }
     }

      socket.onerror = (error) => {
        console.log(error)
        alert(error)
      }

      socket.onclose = ()=> {
        console.log("reconecting")
        setTimeout(() => {
            initWebSocket()
        }, 3000);
      }
}
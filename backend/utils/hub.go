package forum

// import "github.com/gorilla/websocket"

// // Client هو "البورطري" ديال كل مستخدم كونيكطي
// type Client struct {
// 	ID   string          // الـ ID ديالو (othmane1)
// 	Conn *websocket.Conn // السلك (WebSocket) ديالو
// }

// // Hub هو "السنترال" اللي كيسير كولشي
// type Hub struct {
// 	// الماب فين غانقيدو: [السمية] -> [السلك]
// 	Clients map[string]*websocket.Conn

// 	// قناة (Channel) كيدوزو فيها الميساجات اللي بغينا نفرقوهم
// 	Broadcast chan Message

// 	// قناة باش نزيدو ناس جداد (Register)
// 	Register chan *Client

// 	// قناة باش نحيدو الناس اللي خرجو (Unregister)
// 	Unregister chan *Client
// }

// type Message struct {
// 	// النوع: identify, new_message, typing, user_status
// 	Type string `json:"type"`

// 	// المعرفات (ID)
// 	UserID     string `json:"user_id,omitempty"`     // كنستعملوها فـ identify
// 	SenderID   string `json:"sender_id,omitempty"`   // شكون صيفط
// 	ReceiverID string `json:"receiver_id,omitempty"` // لمن غادة

// 	// المحتوى
// 	Content string `json:"content"`

// 	// الوقت (باش نرتبو الشات فـ الداتابيز من بعد)
// 	Timestamp string `json:"timestamp,omitempty"`
// }

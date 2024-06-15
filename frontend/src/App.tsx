import wailsLogo from './assets/wails.png'
import './App.css'
import MySQLConnectionForm from './MySQLConnectionForm'

function App() {
    return (
        <div className="min-h-screen bg-white grid grid-cols-1 place-items-center justify-items-center mx-auto py-8">
            <MySQLConnectionForm />
        </div>
    )
}

export default App

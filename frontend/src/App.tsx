import wailsLogo from './assets/wails.png'
import './App.css'
import MySQLConnectionForm from './MySQLConnectionForm'
import { useState } from 'react';
import DataDisplayComponent from './DataDisplayComponent';

function App() {

    return (
        <div className="min-h-screen bg-white grid grid-cols-1 place-items-center justify-items-center mx-auto py-8">
            <MySQLConnectionForm />
        </div>
    )
}

export default App

import React from "react"
import {NavLink} from 'react-router-dom'

const logOut = () => {
    try {
        // var cookieValue = document.cookie.replace(/(?:(?:^|.*;\s*)csrf_token\s*\=\s*([^;]*).*$)|^.*$/, "$1");
        postNoData("/signout")
        window.location.href = "/"
    } catch (e) {
        alert(e.message);
    }
  }

const postNoData = (url: string) => {
    fetch(url, {
        method: 'POST',
    }).then(res => res)
    .then(response => console.log('Success:', response))
    .catch(error => console.error('Error:', error));
};

export const GlobalHeader: React.StatelessComponent= () => (
    <header className="GlobalHeader">
        <h1>Diary</h1>
        <nav>
            <ul>
                <li><NavLink to="/">Top</NavLink></li>
                <li><NavLink to="/me">Me</NavLink></li>
                <li><span onClick={logOut}>Log Out</span></li>
            </ul>
        </nav>
    </header>
);
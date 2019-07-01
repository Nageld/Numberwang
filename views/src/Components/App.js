import React, {Component} from 'react';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import HomePage from './HomePage';
import GamePage from "./GamePage"
// import NavBar from './NavBar';

class App extends Component {
    render() {
        return (
            <Router>
                <div>
                    <Route name="home" exact path="/" component={HomePage}/>
                    <GamePage/>
                </div>
            </Router>
        )
    }
}

export default App;
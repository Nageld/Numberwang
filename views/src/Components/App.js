import React, {Component} from 'react';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import HomePage from './HomePage';
import NavBar from './NavBar';

class App extends Component {
    render() {
        return (
            <Router>
                <div>
                    <NavBar/>
                    <Route name="home" exact path="/" component={HomePage}/>
                </div>
            </Router>
        )
    }
}

export default App;
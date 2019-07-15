import React, {Component} from 'react';

class TextBox extends Component {
    state = {  }
    render() { 
        return ( 
            <div class="input-group mb-3">
                <input type="text" class="form-control" placeholder="Have a gander at thee!" aria-label="Recipient's username" aria-describedby="basic-addon2"/>
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button">NUMBERWANG?!</button>
                </div>
            </div>
        );
    }
}
 
export default TextBox;
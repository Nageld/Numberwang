import React, {Component} from 'react';
import InputBox from "./InputBox"

const tempSpinnyWheelThingy = {
    width:"200px",
    height:"200px",
    background:"#645FFF",
    borderRadius: "100px",
};

const centeredItems = {
    display:"flex",
    alignItems:'center', 
    justifyContent:'center',

}

class GamePage extends Component {

    render() {
        return(
                <div class="container" style={{marginTop:"60px"}}>
                    <div class="jumbotron text-center">
                        <h1> Lobby (insert num here)</h1>
                        <hr class="my-4"/>
                        
                        <div style={centeredItems}>
                            <div style={tempSpinnyWheelThingy}> (Insert SPINNY WHEEL THINGY here)</div>
                        </div>

                        <div class="row mt-4">
                            <div class="text-center col my-auto">
                                <div class="bg-info">PLAYER 1</div>
                            </div>
                            <div class="text-center col my-auto">
                                <div class="bg-info">PLAYER 2</div>
                            </div>
                            <div class="text-center col my-auto">
                                <div class="bg-info">PLAYER 3</div>
                            </div>
                            <div class="text-center col my-auto">
                                <div class="bg-info">PLAYER 4</div>
                            </div>
                            <div class="text-center col my-auto">
                                <div class="bg-info">PLAYER 5</div>
                            </div>
                        </div>

                        <div class="mt-4">
                            <InputBox />
                            <div> INSERT CHAT BOX HERE </div>
                        </div>


                    </div>
                </div>
            )
    }
}

export default GamePage;
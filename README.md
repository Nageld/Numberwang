# Instructions:
1. Install docker
2. Install docker-compose
3. If running Windows setup WSL to connect to docker
4. run `make default`
5. run `docker-compose up`
6. If running Ubuntu:
	* The container will have working hot reloading
	* After saving a .go file the container will rebuild the app 
7. If not Ubuntu:
	*  Stop and start the container after any changes to reload.

# Functional Requirements:

1.  User Connection
    
	   * P2P
		    * UDP hole punching
		    *  WebRTC
	* Web Server
		* Golang will be used to create web sockets to connect users
	* Game - Numberwang
		* Enter number into text field
		* Display if number is “numberwang” or not
    

  

# Non-Functional Requirements:

1.  Aesthetics  /  Gameplay  Elements

	*  React  used  to  mimic  game

		*  Animations  /  Stylized  components  /  UXUI

2. Train  a  NN  on  numberwang

  

# Goals:

1.  Centralized PoC
    
2.  Setup p2p / centralized web server
    
3.  Make components for the game to utilize
    
4.  Compose components into final game
    
5.  Integrate game into p2p / centralized web server


# Instructions:
How to run the project will go here

# Current State:
1. Basic Go app that has access to network requests
	* MUST be in root of repo, unsure on file hierarchy 
	* Can be deployed to heroku
	* Changes go in main.go
	

![Praise be our mascot/supervisor](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Yellow-headed_caracara_%28Milvago_chimachima%29_on_capybara_%28Hydrochoeris_hydrochaeris%29.JPG/1280px-Yellow-headed_caracara_%28Milvago_chimachima%29_on_capybara_%28Hydrochoeris_hydrochaeris%29.JPG "Praise be our mascot/supervisor")

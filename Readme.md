## Vocal Consonant Counter
Count the number of Vowel and Consonant characters from user input string

## Docker Command : 
  - docker build -t vocalconsonant .
  - docker run -d -p 8080:8080 vocalconsonant

## Usage :
Request to url with params : the string input for the vocal consonant counting function
		
Example: localhost:8080/vocalconsonant/osama

### Note : 
In some case localhost:8080 isn't accessible, instead the app runs and accessible on docker-machine ip

Run command on terminal : docker-machine ip

Example : http://192.168.99.100:8080/vocalconsonant/osama

function addUser(){
    // data to be sent to the POST request
    var _data = {
        username : document.getElementById("username").value,
        email : document.getElementById("email").value,
        userpassword : document.getElementById("userpassword").value
        
    }
    console.log(_data.username)
    fetch('/user/register', {
        method: "POST",
        body: JSON.stringify(_data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    })  
}


// login
var __data = null;
function getAdmin(){
    __data = {
        email : document.getElementById("email").value,
        userpassword : document.getElementById("userpassword").value
    }
    var aname = __data.email
    fetch('/user/login/'+aname)
    .then(res => res.text())
    .then(data => a(data))
}
async function a(credentials){
    const data = JSON.parse(credentials)
    var de = data.Email
    var _de = __data.email
    var dp = data.UserPassword
    var _dp = __data.userpassword
    console.log("dp",dp)
    console.log("_dp",_dp)
    if(de.trim() == _de.trim() && dp.trim() == _dp.trim()){
        location.href = 'data.html'
    }else{
        alert("check the credentials")
    }
    __data = null;
}

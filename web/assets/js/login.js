
const LOGIN_TIME = 30*24*60*60*1000; // 30 days
    
$(function(){
    $("#login").on('click', login);
    initData();
})

function initData(){
    if(!getCookie("SID")){
        // redirect to login
    }else{
        // redirect to target page
    }
}


function login(){
    var username = $("#user").val().trim();
    var password = $("#password").val().trim();
    if (username.length == 0 || password.length == 0) {
        alert("empty username or password ")
    }

    // ajax, set cookies 
    $("#loginbtn").click(function(){
        $.post("/login",
        {
          user: username,
          password: password
        },
        function(data, status){
          alert("Data: " + data + "\nStatus: " + status);
        });
    });
    var sid = data.sessionId; 

    var date = new Date();
    date.setTime(date.getTime() + LOGIN_TIME); 
    document.cookie = "USERNAME=" + Base64.encode(username) + ";path=/;expires="+date.toGMTString();
    document.cookie = "SID=" + Base64.encode(sid) + ";path=/;expires="+date.toGMTString();
    // redirect
}


// logout delete cookies
function logout(){
    var date = new Date();
    date.setTime(date.getTime() - 60*60*1000); 
    document.cookie = "USERNAME;path=/;expires="+date.toGMTString();
    document.cookie = "SID=;path=/;expires="+date.toGMTString();
    // redirect
}

function getCookie(cookieName){
    cookieName += "=";
    var cookieList = document.cookie.split(";");
    for(var i= 0; i< cookieList.length; i++){
        var cookieItem = cookieList[i].trim();
        if(cookieItem.indexOf(cookieName) != -1){
            return cookieItem.substring(cookieName.length, cookieItem.length);
        }
    }
    return '';
}


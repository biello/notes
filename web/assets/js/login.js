
const LOGIN_TIME = 30*24*60*60*1000; // 30 days
    
$(function(){
    initData();
    $("#loginbtn").on('click', login);
    
})

function initData(){
    var sid = getCookie("SID");
    console.log(sid);
    if(sid.length != 0){
        // check session
        checkSession(sid);
    }
}

function checkSession(sid){
    var target = $("#target").val().trim();
    // ajax, set cookies 
    $.post("/signCheck",
    {
        target: target
    },
    function(data, status){
        alert("Data: " + JSON.stringify(data) + "\nStatus: " + status);
        // redirect
        window.location.href = target;
        
    }); 
}


function login(){
    var username = $("#user").val().trim();
    var password = $("#password").val().trim();
    var target = $("#target").val().trim();
    if (username.length == 0 || password.length == 0) {
        alert("empty username or password ")
        return
    }

    // ajax, set cookies 
    $.post("/login",
    {
        user: username,
        password: password,
        target: target
    },
    function(data, status){
        console.log("Data: " + JSON.stringify(data) + "\nHTTP status: " + status);
        if (data.ErrNo == "0") {
            var sid = data.SessionId; 
            var date = new Date();
            date.setTime(date.getTime() + LOGIN_TIME);
            // document.cookie = "USERNAME=" + Base64.encode(username) + ";path=/;expires="+date.toGMTString();
            document.cookie = "SID=" + sid + ";path=/;expires="+date.toGMTString();
            // redirect
            window.location.href = target;
        } else {
            alert("Data: " + JSON.stringify(data) + "\nStatus: " + status);
        }
        
    }); 
}


// logout delete cookies
function logout(){
    var date = new Date();
    date.setTime(date.getTime() - 60*60*1000); 
    // document.cookie = "USERNAME;path=/;expires="+date.toGMTString();
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


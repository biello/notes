    
$(function(){
    $("#passwdbtn").on('click', changePassword);
    
})


function changePassword(){
    var username = $("#user").val().trim();
    var password = $("#password").val().trim();
    var newPassword1 = $("#newPassword1").val().trim();
    var newPassword2 = $("#newPassword2").val().trim();
    if (username.length == 0 || password.length == 0 || newPassword1.length == 0) {
        alert("empty username or password ")
        return
    }

    if (newPassword1 != newPassword2) {
        alert("two new passwords is not the same")
        return
    }

    // ajax, set cookies 
    $.post("/password",
    {
        User: username,
        Password: password,
        NewPassword: newPassword1
    },
    function(data, status){
        console.log("Data: " + JSON.stringify(data) + "\nHTTP status: " + status);
        if (data.ErrNo == "0") {
            alert("success");
            var date = new Date();
            date.setTime(date.getTime() - 60*60*1000); 
            document.cookie = "SID=;path=/;expires="+date.toGMTString();
        } else {
            alert("Data: " + JSON.stringify(data));
        } 
    }); 
}




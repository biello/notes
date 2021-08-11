    
$(function(){
    $("#passwdbtn").on('click', register);
    
})


function register(){
    var username = $("#user").val().trim();
    var password1 = $("#password1").val().trim();
    var password2 = $("#password2").val().trim();
    
    if (username.length == 0 || password1.length == 0 || password2.length == 0) {
        alert("empty username or password ")
        return
    }

    if (password1 != password2) {
        alert("two new passwords is not the same")
        return
    }

    // ajax
    $.post("/admin/register",
    {
        User: username,
        Password: password1,
    },
    function(data, status){
        console.log("Data: " + JSON.stringify(data) + "\nHTTP status: " + status);
        if (data.ErrNo == "0") {
            alert("success");
        } else {
            alert("Data: " + JSON.stringify(data));
        } 
    }); 
}




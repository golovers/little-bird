function upVote(el, id) {
    var userID = getCookie("_user_id");
    if (userID != "") {
        $.ajax({
            type: "POST",
            url: "/api/v1/articles/"+id+"/vote",
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            success: function (data, status, jqXHR) {
                el.querySelector("#article-vote-count-" +id).innerHTML=data.count;
            },
            error: function (jqXHR, status) {
                alert('Oooops: fail to vote. Please login first!');
            }
        });
    } else {
        $(location).attr('href', getCookie("_login_url"))
    }
}
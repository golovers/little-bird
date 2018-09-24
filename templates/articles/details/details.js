function deleteArticle(id) {
    var result = confirm("Delete this article will lead to delete all relevants votes, comments and your audiences will never see this article again. Do you want to continue?");
    if (result) {
        $.ajax({
            type: "DELETE",
            url: "/api/v1/articles/"+id,
            contentType: "application/json; charset=utf-8",
            success: function (data, status, jqXHR) {
                $(location).attr('href', '/articles');
            },
            error: function (jqXHR, status) {
                alert('Oooops: fail to delete: ' + status.code);
            }
        });
    }
}

function requireLogin() {
    var userID = getCookie("_user_id");
    if (userID == "") {
        $(location).attr('href', getCookie("_login_url"))
    }
}
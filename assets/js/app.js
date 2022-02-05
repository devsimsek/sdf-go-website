function redirect(url) {
    window.location.replace(url)
}

function errorAlert(code) {
    switch (code) {
        case "0xednf404":
            return {
                "error": "Document Does Not Exists",
                "message": "The document that you've searched for does not exists in sdf documentary. \n(Info: documentary is still not complete. Please retry again.)",
            };
        case "0xnr":
            return {
                "error": "Not Ready",
                "message": "Functionality is not yet ready. Please be avare that sdf website is under development.",
            };
    }
}
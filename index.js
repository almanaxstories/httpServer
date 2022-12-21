console.log('JS-file connected successfully!');

window.addEventListener('load', ()=>{
let appContainer = document.getElementById('signupFormDiv');
let signupForm = newElement('FORM');
signupForm.method = 'POST';
appContainer.append(signupForm);
let firstNameLabel = newElement("label", "", "First Name:", "fNameText", "fNameLabel");
signupForm.append(firstNameLabel);
let firstNameInput = newElement("input", "text", "", "", "fNameText");
signupForm.append(firstNameInput);
let lastNameLabel = newElement("label", "", "Last Name:", "lNameText", "lNameLabel");
signupForm.append(lastNameLabel);
let lastNameInput = newElement("input", "text", "", "", "lNameText");
signupForm.append(lastNameInput);
let submitButton = newElement("input", "submit", "Let's go!", "", "submitButton");
signupForm.append(submitButton);

})


function newElement(tag, typeOfInput, text, labelFor, id){
    let element = document.createElement(tag);

    if (typeOfInput){
        element.type = typeOfInput;
    }
    
    if (text){
        element.innerText = text;
    }
    
    if (labelFor){
        element.for = labelFor;
    }

    if (id){
        element.id = id;
    }    
    return element;
}
package mydict

import (
    "errors"
)

// Dictionary type
type Dictionary map[string]string

//  Create methods for dictionary

var (
    errNotFound = errors.New("Not Found")
    errWordExists = errors.New("That word already exists")
    errCantUpdate = errors.New("Can't update non-existing word")
    errCantDelete = errors.New("Can't delete non-existing word")
)

//  Search for a word
func (d Dictionary) Search(word string) (string, error) {
    value, exists := d[word]   //  string and bool
    if exists {
        return value, nil
    }
//    return errors.New("Not found")
    return "", errNotFound
}

//  Add new word
func (d Dictionary) Add(word, def string) error {
    //  return error if word does not exists
    _, err := d.Search(word)

    /*
    if err == errNotFound {
        d[word] = def
    } else if err == nil{
        return errWordExists
    }
    return nil
    */

    switch err{
    case errNotFound:
        d[word] = def
    case nil:
        return errWordExists
    }
    return nil
}

//  Update a word
func (d Dictionary) Update (word, definition string) error {
    _, err := d.Search(word)
    switch err {
    case nil:
        d[word] = definition
    case errNotFound:
        return errCantUpdate
    }
    return nil
}

//  Delete an existing word
func (d Dictionary) Delete (word string) error{
    _, err := d.Search(word)
    
    switch err{
    case nil:
        delete(d, word)
    case errNotFound:
        return errCantDelete
    }
    return nil
}

package com.sakrafux.sem.realworld.exception.response;

public class NotFoundResponseException extends Exception {

    public NotFoundResponseException() {
        super();
    }

    public NotFoundResponseException(String message) {
        super(message);
    }

    public NotFoundResponseException(String message, Throwable cause) {
        super(message, cause);
    }

    public NotFoundResponseException(Throwable cause) {
        super(cause);
    }
}

package com.sakrafux.sem.realworld.exception.response;

public class GenericErrorResponseException extends Exception {

    public GenericErrorResponseException() {
        super();
    }

    public GenericErrorResponseException(String message) {
        super(message);
    }

    public GenericErrorResponseException(String message, Throwable cause) {
        super(message, cause);
    }

    public GenericErrorResponseException(Throwable cause) {
        super(cause);
    }
}

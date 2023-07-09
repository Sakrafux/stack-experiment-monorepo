package com.sakrafux.sem.realworld.exception.response;

public class UnauthorizedResponseException extends Exception {

    public UnauthorizedResponseException() {
        super();
    }

    public UnauthorizedResponseException(String message) {
        super(message);
    }

    public UnauthorizedResponseException(String message, Throwable cause) {
        super(message, cause);
    }

    public UnauthorizedResponseException(Throwable cause) {
        super(cause);
    }
}

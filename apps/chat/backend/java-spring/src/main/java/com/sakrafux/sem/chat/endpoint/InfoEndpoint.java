package com.sakrafux.sem.chat.endpoint;

import lombok.RequiredArgsConstructor;
import org.springframework.security.access.annotation.Secured;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/info")
@RequiredArgsConstructor
public class InfoEndpoint {

    @Secured("ROLE_USER")
    @GetMapping
    public String getInfo() {
        return "This is the chat service";
    }

}

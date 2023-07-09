package com.sakrafux.sem.realworld.logging;

import java.util.Arrays;
import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Aspect
@Component
@Slf4j
public class LoggingAspect {

    @Pointcut("within(com.sakrafux.sem.realworld.endpoint..*)")
    public void endpointPointcut() {
    }

    @Pointcut("within(com.sakrafux.sem.realworld.service..*)")
    public void servicePointcut() {
    }

    @Pointcut("within(@org.springframework.stereotype.Repository *)")
    public void repositoryPointcut() {
    }

    @Around("servicePointcut() || repositoryPointcut()")
    public Object logMethod(ProceedingJoinPoint joinPoint) throws Throwable {
        long startTime = System.currentTimeMillis();

        log.info(
            "Before method: " + joinPoint.getSignature().getDeclaringType().getSimpleName() + "."
                + joinPoint.getSignature().getName());
        log.debug("Called with argument(s): " + getJoinPointArgs(joinPoint));

        Object result = joinPoint.proceed();

        long endTime = System.currentTimeMillis();
        long duration = endTime - startTime;
        log.info("After method in " + duration + "ms: " +
            joinPoint.getSignature().getDeclaringType().getSimpleName() + "."
            + joinPoint.getSignature().getName());
        log.debug("With result: " + (result != null ? result.toString() : "null"));
        return result;
    }

    @Around("endpointPointcut()")
    public Object logEndpoint(ProceedingJoinPoint joinPoint) throws Throwable {
        long startTime = System.currentTimeMillis();

        var annotations = ((MethodSignature) joinPoint.getSignature()).getMethod().getAnnotations();

        String[] path = {};
        String method = "";
        for (var annotation : annotations) {
            if (annotation instanceof GetMapping getMapping) {
                path = getMapping.value();
                method = "GET";
            }
            if (annotation instanceof PostMapping postMapping) {
                path = postMapping.value();
                method = "POST";
            }
            if (annotation instanceof PutMapping putMapping) {
                path = putMapping.value();
                method = "PUT";
            }
            if (annotation instanceof DeleteMapping deleteMapping) {
                path = deleteMapping.value();
                method = "DELETE";
            }
        }

        var classAnnotations = joinPoint.getSignature().getDeclaringType().getAnnotations();
        String[] rootPath = {};
        for (var annotation : classAnnotations) {
            if (annotation instanceof RequestMapping requestMapping) {
                rootPath = requestMapping.value();
            }
        }

        if (rootPath.length == 0) {
            rootPath = new String[]{""};
        }
        if (path.length == 0) {
            path = new String[]{""};
        }

        log.info(method + "-Endpoint: " + rootPath[0] + path[0]);
        log.info(
            "Before method: " + joinPoint.getSignature().getDeclaringType().getSimpleName() + "."
                + joinPoint.getSignature().getName());
        log.debug("Called with argument(s): " + getJoinPointArgs(joinPoint));

        Object result = joinPoint.proceed();

        long endTime = System.currentTimeMillis();
        long duration = endTime - startTime;
        log.info(method + "-Endpoint: " + rootPath[0] + path[0]);
        log.info("After method in " + duration + "ms: " +
            joinPoint.getSignature().getDeclaringType().getSimpleName() + "."
            + joinPoint.getSignature().getName());
        log.debug("With result: " + (result != null ? result.toString() : "null"));
        return result;
    }

    private String getJoinPointArgs(ProceedingJoinPoint joinPoint) {
        var params = ((MethodSignature) joinPoint.getSignature()).getParameterNames();
        var args = joinPoint.getArgs();

        var sb = new StringBuilder();

        for (int i = 0; i < args.length; i++) {
            if (i > 0) {
                sb.append(", ");
            }
            sb.append(params != null ? params[i] : "");
            sb.append("=");
            sb.append(args[i] != null ? args[i].toString() : "null");
        }

        return sb.toString();
    }

}

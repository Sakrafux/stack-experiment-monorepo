package com.sakrafux.sem.realworld.dto.request;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class PaginationParamDto {

    @Min(0)
    private int offset = 0;

    @Min(1)
    @Max(20)
    private int limit = 1;

}

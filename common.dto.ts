import { ApiProperty } from "@nestjs/swagger";
import { IsNumber, IsOptional } from "class-validator";


export class PaginationDto {

    @ApiProperty({ description: 'The number of results to skip. It is optional and should be a number.', required: false })
    @IsOptional()
    @IsNumber()
    skipResults?: number;

    @ApiProperty({ description: 'The number of results to return. It is optional and should be a number.', required: false })
    @IsOptional()
    @IsNumber()
    takeResultAmount?: number;
}
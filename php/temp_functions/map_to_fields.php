<?php


/**
 * 把请求参数映射到相关的数据表字段, `key`值为请求参数的`key`，`value`为映射到的字段，例如：
 * `child_name`是请求的参数，`student_name`是数据表的字段；
 * $requestParams = [
 *     'guardian_id' => 25,
 *     'child_name' => '小蓝',
 *     'idcard' => '420606199810111234',
 *     'institute_id' => 1                
 * ];
 * $studentFieldsMap = [
 *     'child_name' => 'student_name',
 *     'idcard'
 * ];
 * $studentFields = [];
 * 
 * @param array $requestParams 请求参数数组；使用 TP5 框架封装的`request`类； 
 * @param array $fieldsMap 数据表字段的映射数组；
 * @return array $fields 映射完成后的数据表字段数组；
 */
function mapToFields($requestParams, $fieldsMap)
{
    $fields = [];
    foreach ($fieldsMap as $mapKey => $mapValue) {
        foreach ($requestParams as $key => $value) {
            if ($key == $mapKey) {
                $fields[$mapValue] = $requestParams[$key];
                continue;
            }
            if ($mapValue == $key) {
                echo 'map_value: ' . $mapValue . '<br />';
                $fields[$value] = $requestParams[$key];
                continue;
            }
        }
    }
    return $fields;
}

$requestParams = [
    'guardian_id' => 25,
    'child_name' => '小蓝',
    'idcard' => '420606199810111234',
    'institute_id' => 1                
];
$studentFieldsMap = [
    'child_name' => 'student_name',
    'idcard'
];
$studentFields = mapToFields($requestParams, $studentFieldsMap);
var_dump($studentFields);

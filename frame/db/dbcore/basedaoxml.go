package dbcore

const BaseXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>
	<insert id="Save">
			insert into {{.TableName}}(
			{{range $index, $ele := $.TableColumn}}{{if $index}},{{end}}{{$ele.ColumnName}}{{end}}
			) values (
			{{range $index, $ele := $.TableColumn}}{{if $index}},{{end}}#{{{$ele.FieldName}}}{{end}}
			)
	</insert>
</mapper>

`

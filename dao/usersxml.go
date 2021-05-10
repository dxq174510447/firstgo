package dao

const UsersXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>
	<select id="GetById">
			select * from users where id = #{id}
	</select>
	
	<select id="FindByNameAndStatus">
			select * from users where   
			<include refid="conditionA1"></include>
           order by id desc 
	</select>

	<sql id="conditionA1">
		name = #{name} and status = #{status} 
		{{if .statusList}}
			and status in (
				{{range $index, $ele := $.statusList}}{{if $index}},{{end}}#{statusList[{{$index}}]}{{end}}
			)
		{{end}}
	</sql>

	<select id="FindIds">
			select id from users order by id desc 
	</select>

	<select id="FindNames">
			select distinct name from users 
			where 1=1
<![CDATA[
			{{if .NameIn}}
				and name in (
					{{range $index, $ele := $.NameIn}}{{if $index}},{{end}}#{NameIn[{{$index}}]}{{end}}
				)
			{{end}}
			and status < 1
]]>
			order by id desc 
	</select>

	<select id="FindFees">
			select distinct fee from users  order by id desc 
	</select>
	
	<select id="GetMaxFees">
			select max(fee)  from users 
	</select>

	<select id="GetMaxId">
			select max(id)  from users 
	</select>

	<update id="UpdateNameByField">
		update users set name = #{name} where id = #{id}
	</update>

	<update id="UpdateNameByEntity">
		update users set name = #{Name} where id = #{Id}
	</update>

	<delete id="DeleteNameByField">
		delete from users where name = #{name} and id = #{id}
	</delete>

	<delete id="DeleteNameByEntity">
		delete from users where name = #{Name} and id = #{Id} 
		{{if .NameIn}}
			and name in (
				{{range $index, $ele := $.NameIn}}{{if $index}},{{end}}#{NameIn[{{$index}}]}{{end}}
			)
		{{end}}
	</delete>

	<insert id="InsertSingle">
		insert into users(name,password) values (#{Name},#{Password}) 
	</insert>

	<insert id="InsertBatch">
		insert into users(name,password) values 
		{{range $index, $ele := .values}}
		{{if $index}},{{end}}(#{values[{{$index}}].Name},#{values[{{$index}}].Password})
		{{end}}
	</insert>
	
</mapper>

`

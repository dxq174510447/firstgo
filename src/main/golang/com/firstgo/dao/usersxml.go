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
		insert into users(name,password,status,fee,create_date,create_time,fee_total) values 
			(#{Name},#{Password},#{Status},#{Fee},#{CreateDate},#{CreateTime},#{FeeTotal}) 
	</insert>

	<insert id="InsertBatch">
		insert into users(name,password) values 
		{{range $index, $ele := .values}}
		{{if $index}},{{end}}(#{values[{{$index}}].Name},#{values[{{$index}}].Password})
		{{end}}
	</insert>


	<insert id="Save1">
		insert into users(name,password,status,fee,create_date,create_time,fee_total) values 
			(#{Name},#{Password},#{Status},#{Fee},#{CreateDate},#{CreateTime},#{FeeTotal})  
	</insert>

	<update id="Update1">
		update users set name = #{Name},status = #{Status},password = #{Password} where id = #{Id} 
	</update>

	<delete id="Delete1">
		delete from users where id = #{id}
	</delete>
	
	<select id="Get1">
			select * from users where id = #{id}
	</select>

	<update id="ChangeStatus1">
		update users set status = #{status} where id = #{id} 
	</update>

	<select id="List1">
		select * from users limit #{Page},#{PageSize}
	</select>

	<select id="FindByNameExcludeId1">
		select count(*) from users where name = #{name} and id != #{id}
	</select>

	<select id="FindByName1">
		select count(*) from users where name = #{name}
	</select>

	<select id="QueryAddon">
		select a.id_,a.company_id_ from base_business a,base_business_add b
			where a.id_ = b.BUSINESS_ID_
			and a.DEPARTMENT_ID_ = 1
			{{if .CheckType}}
			and a.STATUS_ = #{CheckType}
			{{end}}
			{{if .AddonTitle}}
			and a.NAME_ like #{AddonTitle}
			{{end}}
			limit 0,10
	</select>
</mapper>

`

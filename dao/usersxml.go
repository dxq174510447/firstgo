package dao

const UsersXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>
	<select id="GetById">
			select * from users where id = #{id}
	</select>
	
	<select id="FindByNameAndStatus">
			select * from users where name = #{name} and status = #{status}  order by id desc 
	</select>

	<select id="FindIds">
			select id from users order by id desc 
	</select>

	<select id="FindNames">
			select distinct name from users order by id desc 
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

	
</mapper>

`

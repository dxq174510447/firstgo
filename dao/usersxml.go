package dao

const UsersXml = `

<?xml version="1.0" encoding="UTF-8" ?>
<mapper>

	<sql id="pageImport">123
    xxxxxx
	</sql>

	<select id="page1">
		select * from a where 
		<include refid="pageImport"></include>
	</select>

	<update id="xxx1">
a
    </update>

	<delete id="xxx2">
b
    </delete>

	<insert id="xxx3">
c
    </insert>
</mapper>

`
